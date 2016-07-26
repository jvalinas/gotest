package udpproto

import (
  "artifact"
  "encoding/gob"
  "bytes"
  "net"
  "logging"
)

const (
  port = "10001"
  pkgLimit = 5
)

type GamePkg struct {
  ServerName string
  Artifacts map[int]*artifact.Artifact
}

type EventPkg struct {
  Name int
  Artifacts map[int]*artifact.Artifact
}

type dataPkg struct {
  name string
  slot int
  ip string
  pkgType string
  data [][]byte
}

type netDataPkg struct {
  Name string
  PkgType string
  Length int
  Data [][]byte
}

func NewGamePkg(serverName string, artifacts map[int]*artifact.Artifact) GamePkg {
  return GamePkg{ServerName: serverName, Artifacts: artifacts}
}

func newDataPkg(name string, slot int, pkgType string, data [][]byte)  dataPkg {
  return dataPkg{name: name, slot: slot, pkgType: pkgType, data: data}
}

func newDataPkgFomNetPkg(netDataPkg netDataPkg) *dataPkg {
  //logging.Printf("Received net data part from %s: %v",netDataPkg.Name, netDataPkg)
  //logging.Printf("Received net data part from %s: Total %s, Re %s",netDataPkg.Name, netDataPkg.Length, len(netDataPkg.Data))
  slot := decodeInt(netDataPkg.Data[0])
  pkg := newDataPkg(netDataPkg.Name, slot, netDataPkg.PkgType, netDataPkg.Data[1:])
  return &pkg
}

func (dp dataPkg) len() int {
    return len(dp.data) + 1
}

func newNetDataPkgFromBytes(data []byte) netDataPkg {
  d := gob.NewDecoder(bytes.NewBuffer(data))

  // Decoding the serialized data
  var netDataPkg netDataPkg
  err := d.Decode(&netDataPkg)
  checkError(err)
  return netDataPkg
}

func (np netDataPkg) Bytes() []byte {
  data := new(bytes.Buffer)
  e := gob.NewEncoder(data)

  // Encoding
  err := e.Encode(np)
  checkError(err)
  return data.Bytes()
}

func Start(serverIp string, in chan GamePkg, out chan GamePkg, event chan EventPkg) {
  //serverIp := "172.17.0.1:10001"
  localAddr := ":" + port
  receivQueue := make(chan dataPkg, 100)
  sendQueue := make(chan dataPkg, 100)
  go serverStart(localAddr, receivQueue)
  go clientStart(serverIp, sendQueue)
  go proccessData(in, out, event, receivQueue, sendQueue)
}

func proccessData(in chan GamePkg, out chan GamePkg, event chan EventPkg,
    receivQueue chan dataPkg, sendQueue chan dataPkg) {
  for {
    select {
      case data := <- receivQueue:
        if data.pkgType == "gamedata" {
          logging.Println("Receiving new package from: ",
                          data.name, " with ",
                          len(data.data), " artifacts.")
	  t1 := pkgToGameData(data)
          for _, a := range t1.Artifacts {
          logging.Println("Received art: ", a.Pos() )
}
          out <- t1
        } else {
          logging.Println("Unknown package type: ", data.pkgType)
        }
      case data := <- in:
        //logging.Println("Sending new package from: ",
        //                data.ServerName, " with ",
        //                len(data.Artifacts), " artifacts.")
        sendQueue <- gameDataToPkg(data)
    }
  }
}

func gameDataToPkg(gp GamePkg) dataPkg {
  pkgType := "gamedata"
  data:= make([][]byte, 0, len(gp.Artifacts))
  for _, artifact := range gp.Artifacts {
    //logging.Println("Appending data? ", artifact.Bytes())
    data = append(data, artifact.Bytes())
  }
  //logging.Println("gameToPkg: Artifacts: ", len(gp.Artifacts)," Data to send:", len(data))
  return newDataPkg(gp.ServerName, 0, pkgType, data)
}

func pkgToGameData(dp dataPkg) GamePkg {
  artifacts := make(map[int]*artifact.Artifact, len(dp.data))
  //logging.Printf("Processing data from %s: %v",dp.name, dp)
  for _, chunk := range dp.data {
    artifact := artifact.NewActifactFromBytes(chunk)
    artifacts[artifact.Id()] = &artifact
  }
  return GamePkg{ServerName: dp.name, Artifacts: artifacts}
}

func checkError(err error) {
    if err  != nil {
        logging.Println("Error: " , err)
    }
}

func serverStart(listenIp string, queue chan dataPkg) {
    /* Lets prepare a address at any address at port 10001*/
    ServerAddr,err := net.ResolveUDPAddr("udp", listenIp)
    checkError(err)

    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    checkError(err)
    logging.Println("Server listening at: ",  listenIp)
    defer ServerConn.Close()
    pkgs := make(map[string]*dataPkg)
    for {
      _, newPkg := receivePkg(ServerConn)

      pkgKey := newPkg.Name + newPkg.PkgType

      var pkg *dataPkg
      if _, ok := pkgs[pkgKey]; !ok {
        pkg = newDataPkgFomNetPkg(newPkg)
        //logging.Println("New NetPkg from slot: ", newPkg.Name)
        pkgs[pkgKey] = pkg
      } else {
        pkg = pkgs[pkgKey]
        //logging.Println("Received ", len(newPkg.Data),
        //                " data from slot", newPkg.Name)
        for _, data := range newPkg.Data {
          pkg.data = append(pkg.data, data)
        }
      }
      //logging.Printf("New net package: %s of %s", len(pkg.data), newPkg.Length)
      if pkg.len() == newPkg.Length {
        //logging.Printf("Received net data from %s: %v",pkg.name, pkg)
        queue <- *pkg
        delete(pkgs, pkgKey)
      }
    }
}

func clientStart(serverIp string, queue chan dataPkg) {
    ServerAddr,err := net.ResolveUDPAddr("udp", serverIp)
    checkError(err)

    LocalAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:0")
    checkError(err)

    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    checkError(err)
    logging.Println("Client connected to: ",  serverIp)
    defer Conn.Close()
    for {
        dataPkg := <- queue
        sendNetPackage(Conn, dataPkg)
    }
}

func sendNetPackage(Conn *net.UDPConn, dataPkg dataPkg)  {
  var netPkg netDataPkg
  var ctam int
  length := len(dataPkg.data) + 1
  //numChunks := length/pkgLimit
  //logging.Println("data to send: ", len(dataPkg.data))
  if ctam = pkgLimit; length < pkgLimit {
    ctam = length
  }
  chunks := make([][]byte, 0, ctam)
  chunks = append(chunks, encodeInt(dataPkg.slot))
  i := 2
  for _, chunk := range dataPkg.data {
    if i % pkgLimit == 0 {
      netPkg = netDataPkg{dataPkg.name, dataPkg.pkgType, length, chunks}
      sendBytes(Conn, netPkg.Bytes())
      if ctam = pkgLimit; i - length < pkgLimit {
        ctam = i - length
      }
      chunks = make([][]byte, 0, ctam)
    }
    chunks = append(chunks, chunk)
    //logging.Println("Adding chunk: ", chunk)
    i++
  }
  if len(chunks) != 0 {
    netPkg = netDataPkg{dataPkg.name, dataPkg.pkgType, length, chunks}
    //logging.Printf("Sending net data part from %s: %v",netPkg.Name, netPkg)
    //logging.Printf("Sending from %s: t: %s p: %s",netPkg.Name, netPkg.Length, len(netPkg.Data))
    sendBytes(Conn, netPkg.Bytes())
  }
}

func receivePkg(ServerConn *net.UDPConn) (string, netDataPkg) {
    var newPkg netDataPkg
    data := readBytes(ServerConn)
    if data != nil {
      newPkg = newNetDataPkgFromBytes(data)
      //logging.Println(" Received from :", newPkg)
      return "", newPkg
    }
    return "Non Data Received", newPkg
  }

func readBytes(ServerConn *net.UDPConn) []byte {
  buf := make([]byte, 1024)
  n, addr, err := ServerConn.ReadFromUDP(buf)
  if err != nil {
      logging.Println("Error: ",err, " from: ", addr)
  } else {
      return buf[0:n]
  }
  return nil
}

func sendBytes(Conn *net.UDPConn, data []byte) {
  _,err := Conn.Write(data)
  if err != nil {
      logging.Println(err)
  }
}

func encodeInt(val int) []byte {
  data := new(bytes.Buffer)
  e := gob.NewEncoder(data)
  err := e.Encode(val)
  if err != nil {
      logging.Println("NetPkgData encode failure: ", err)
  }
  return data.Bytes()
}

func encodeStr(val string) []byte {
  data := new(bytes.Buffer)
  e := gob.NewEncoder(data)
  err := e.Encode(val)
  if err != nil {
      logging.Println("NetPkgData encode failure: ", err)
  }
  return data.Bytes()
}

func decodeInt(data []byte) int {
  d := gob.NewDecoder(bytes.NewBuffer(data))
  var val int
  err := d.Decode(&val)
  checkError(err)
  return val
}

func decodeStr(data []byte) string {
  d := gob.NewDecoder(bytes.NewBuffer(data))
  var val string
  err := d.Decode(&val)
  checkError(err)
  return val
}

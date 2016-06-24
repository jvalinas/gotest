package network

import (
  "artifact"
  _"fmt"
  "net"
  _"time"
  "netinfo"
)

func checkError(err error) {
    if err  != nil {
        //fmt.Println("Error: " , err)
    }
}

func ServerStart(listenIp string, queue chan netinfo.NetPackage) {
    /* Lets prepare a address at any address at port 10001*/
    ServerAddr,err := net.ResolveUDPAddr("udp",listenIp)
    checkError(err)

    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    checkError(err)
    defer ServerConn.Close()
    netPackages := make(map[int]*netinfo.NetPackage)
    for {
      netPkgData := receiveNetPackageData(ServerConn)

      if _, ok := netPackages[netPkgData.Name]; !ok {
        //fmt.Println("New NetPkg from slot: ", netPkgData.Name)
        artifacts := make(map[int]*artifact.Artifact)
        netPkg := netinfo.NewNetPackage(netPkgData.Name, artifacts)
        netPackages[netPkgData.Name] = &netPkg
      }
      netPkg := netPackages[netPkgData.Name]
      //fmt.Println("Received ", len(netPkgData.Artifacts), " artifacts from slot", netPkgData.Name)
      for _, data := range netPkgData.Artifacts {
        artifact := artifact.NewActifactFromBytes(data)
        netPkg.Artifacts[artifact.Id()] = &artifact
      }
      if len(netPkg.Artifacts) == netPkgData.NumArtifacts {
        queue <- *netPkg
        delete(netPackages, netPkg.Name)
      }
    }
}

func ClientStart(serverIp string, queue chan netinfo.NetPackage) {
    ServerAddr,err := net.ResolveUDPAddr("udp", serverIp)
    checkError(err)

    LocalAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:0")
    checkError(err)

    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    checkError(err)

    defer Conn.Close()
    for {
        netPackage := <-queue
        sendNetPackage(Conn, netPackage)
    }
}

func Start(){
  workqueue := make(chan netinfo.NetPackage, 100)
  go ServerStart(":10001", workqueue)
  go ClientStart("172.17.0.2:10001", workqueue)
}

func sendNetPackage(Conn *net.UDPConn, ni netinfo.NetPackage)  {
  numArtifacts := len(ni.Artifacts)
  i := 1
  numChunks := numArtifacts/5
  chunkOfArtifacts := make([][]byte, 0, numChunks)
  for _, artifact := range ni.Artifacts {
    if i == 5 {
      netPkgData := netinfo.NewNetPackageData(ni.Name, numArtifacts, chunkOfArtifacts)
      sendBytes(Conn, netPkgData.Bytes())
      chunkOfArtifacts = make([][]byte, 0, numChunks)
      i = 1
    }
    chunkOfArtifacts = append(chunkOfArtifacts, artifact.Bytes())
    i++
  }
  if len(chunkOfArtifacts) != 0 {
    netkgData := netinfo.NewNetPackageData(ni.Name, numArtifacts, chunkOfArtifacts)
    sendBytes(Conn, netkgData.Bytes())
  }
}

func receiveNetPackageData(ServerConn *net.UDPConn) *netinfo.NetPkgData {
    data := readBytes(ServerConn)
    if data != nil {
      netPkgData := netinfo.NewPackageDataFromBytes(data)
      //fmt.Println(" Received from :", netPkgData)
      return &netPkgData
    }
    return nil
  }

func readBytes(ServerConn *net.UDPConn) []byte {
  buf := make([]byte, 1024)
  n, _, err := ServerConn.ReadFromUDP(buf)
  if err != nil {
      //fmt.Println("Error: ",err, " from: ", addr)
  } else {
      return buf[0:n]
  }
  return nil
}

func sendBytes(Conn *net.UDPConn, data []byte) {
  _,err := Conn.Write(data)
  if err != nil {
      //fmt.Println(err)
  }
}

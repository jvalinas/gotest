package netinfo

import (
  "artifact"
  _"fmt"
  "encoding/gob"
  "bytes"
)

type NetPackage struct {
  Name int
  Artifacts map[int]*artifact.Artifact
}

type NetPkgData struct {
  Name int
  NumArtifacts int
  Artifacts [][]byte
}

func NewNetPackage(name int, artifacts map[int]*artifact.Artifact)  NetPackage {
  return NetPackage{Name: name, Artifacts: artifacts}
}

func NewNetPackageData(name int, numArtifacts int, artifacts [][]byte) NetPkgData {
  return NetPkgData{Name: name, NumArtifacts: numArtifacts, Artifacts: artifacts}
}

func NewPackageDataFromBytes(data []byte) NetPkgData {
  d := gob.NewDecoder(bytes.NewBuffer(data))

  // Decoding the serialized data
  var netPkgData NetPkgData
  err := d.Decode(&netPkgData)
  if err != nil {
      //panic(err)
      //fmt.Println("netPkgData decode failure: ", err)
  }
  return netPkgData
}

func (p NetPkgData) Bytes() []byte {
  data := new(bytes.Buffer)
  e := gob.NewEncoder(data)

  // Encoding
  err := e.Encode(p)
  if err != nil {
      //fmt.Println("NetPkgData encode failure: ", err)
  }
  return data.Bytes()
}

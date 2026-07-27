package tests

import (

"testing"

"uav-telemetry-visualizer/internal/parser"

)

func TestLoad(t *testing.T){

rows:=parser.Load()

if len(rows)==0{

t.Fail()

}

}

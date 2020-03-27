@echo off
rsrc -manifest elevator.manifest -o ..\cmd\elevator\elevator.syso
pushd ..\cmd\elevator
go install
popd

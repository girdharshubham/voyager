package edu.self.domain

import com.google.protobuf.empty.Empty
import edu.self
import edu.self.{CreateDeviceData, GetDeviceData}
import kalix.scalasdk.valueentity.ValueEntity
import kalix.scalasdk.valueentity.ValueEntityContext

// This class was initially generated based on the .proto definition by Kalix tooling.
//
// As long as this file exists it will not be overwritten: you can maintain it yourself,
// or delete it so it is regenerated as needed.

class Device(context: ValueEntityContext) extends AbstractDevice {
  override def emptyState: DeviceState = DeviceState()

  override def addLocation(currentState: DeviceState, addLocationData: self.AddLocationData): ValueEntity.Effect[Empty] =
    effects
      .updateState(currentState.copy(location = addLocationData.location))
      .thenReply(Empty.defaultInstance)

  override def createDevice(currentState: DeviceState, createDeviceData: CreateDeviceData): ValueEntity.Effect[Empty] =
    effects
      .updateState(DeviceState(createDeviceData.deviceId))
      .thenReply(Empty.defaultInstance)

  override def getDevice(currentState: DeviceState, getDeviceData: GetDeviceData): ValueEntity.Effect[DeviceState] =
    effects.reply(currentState)
}


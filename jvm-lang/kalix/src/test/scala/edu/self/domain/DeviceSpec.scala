package edu.self.domain

import com.google.protobuf.empty.Empty
import edu.self
import kalix.scalasdk.testkit.ValueEntityResult
import kalix.scalasdk.valueentity.ValueEntity
import org.scalatest.matchers.should.Matchers
import org.scalatest.wordspec.AnyWordSpec

class DeviceSpec
    extends AnyWordSpec
    with Matchers {

  "Device" must {

    "have example test that can be removed" in {
      val service = DeviceTestKit(new Device(_))
      pending
      // use the testkit to execute a command
      // and verify final updated state:
      // val result = service.someOperation(SomeRequest)
      // verify the reply
      // val reply = result.getReply()
      // reply shouldBe expectedReply
      // verify the final state after the command
      // service.currentState() shouldBe expectedState
    }

    "handle command AddLocation" in {
      val service = DeviceTestKit(new Device(_))
      pending
      // val result = service.addLocation(self.AddLocationData(...))
    }

  }
}

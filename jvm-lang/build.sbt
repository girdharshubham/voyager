ThisBuild / version := "0.1.0-SNAPSHOT"

lazy val root = (project in file("."))
  .aggregate(kalix)

lazy val kalix = (project in file("kalix"))
  .enablePlugins(KalixPlugin)
  .settings(
    scalaVersion := "2.13.10"
  )

lazy val v3 = (project in file("v3"))
  .settings(
    scalaVersion := "3.2.2"
  )

val scala3Version = "3.8.3"

lazy val root = project
  .in(file("."))
  .settings(
    name := "todolist_example",
    version := "0.1.0-SNAPSHOT",

    scalaVersion := scala3Version,

    libraryDependencies ++= Seq(
      "net.aichler" % "jupiter-interface" % JupiterKeys.jupiterVersion.value % Test // using JUnit 5
    ),
    Test / testFrameworks += new TestFramework("org.junit.platform.sbt.JUnitPlatformFramework")
  )

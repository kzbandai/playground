scalaVersion := "2.12.4"

scalacOptions ++= Seq("-deprecation", "-feature", "-unchecked", "-Xlint")
// deprecation: 今後廃止予定の記述の警告
// feature: 注意しなければならない機能や新機能を利用していることの警告
// unchecked: 型消去などでパターンマッチが有効に機能しない場合
// Xlint: その他、望ましい書き方や落とし穴についての情報
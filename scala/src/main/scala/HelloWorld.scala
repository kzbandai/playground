object HelloWorld {
  var num: Int = 1
  var str: String = "start"
  var short: Short = _ // initial value

  var NUM: Int = 1
  val STR: String = "2"
  var array: Array[String] = Array("a", "b", "c")

  def main(args: Array[String]) = {
    println("Hello World")

    println(num)
    println(str)
    println(short)

    println(NUM)
    println(STR)

    if (true) println(STR) else println("else")

    for (s <- array) println(s)

    val one = 1
    one match {
      case 1 => println("one")
      case 2 => println("two")
      case _ => println("other")
    }

    println(append_ten(NUM))
    println(append_eight(NUM))
  }

  def append_ten(num: Int): Int = {
    num + 10
  }

  def append_eight(num: Int): Int = num + 8
}

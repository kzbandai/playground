<?php

class HelloWorld
{
    public function __invoke(string $name, int $a)
    {
        echo 'Hello ' . $name;
        echo $a;
    }
}

$reflectionMethod = new \ReflectionMethod('HelloWorld', '__invoke');
echo $reflectionMethod->invoke(new HelloWorld(), ...['Mike', 8]);

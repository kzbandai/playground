<?php

interface HelperInterface
{
}

class HogeAndFuga implements HelperInterface
{
    public function __invoke(array $hoge = ['hoge'])
    {
        var_dump($hoge);
    }
}

trait Hoge
{
    public function __call($name, $arguments)
    {
        $strings = mb_split('_', $name);
        $class_name = '';
        foreach ($strings as $s) {
            $class_name .= ucfirst($s);
        }

        if (($c = new $class_name) instanceof HelperInterface) {
            (new $c)($arguments);
        }
    }
}

(new class()
{
    use Hoge;

    public function main()
    {
        $this->hoge_and_fuga('aaa');
    }
})->main();

(new class()
{
    use Hoge;
}
)->hoge_and_fuga('bbb');

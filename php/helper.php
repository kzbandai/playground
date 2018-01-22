<?php

interface HelperInterface
{
    public function __invoke();
}

class HogeAndFuga implements HelperInterface
{
    public function __invoke()
    {
        echo 'hoge';
    }
}

trait Hoge
{
    public function __call($name, $arguments)
    {
        $array = mb_split('_', $name);
        $str = '';
        foreach ($array as $s) {
            $str .= ucfirst($s);
        }

        $do = new $str;
        $do();
    }
}

(new class()
{
    use Hoge;

    public function main()
    {
        $this->hoge_and_fuga();
    }
})->main();

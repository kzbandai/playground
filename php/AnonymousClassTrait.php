<?php


/**
 * @method int hoge
 */
trait AnonymousClassTrait
{
    public function __construct(array $data)
    {
        foreach ($data as $key => $value) {
            if (is_string($key)) {
                $key = strtolower($key);

                $this->$key = $value;
            }
        }
    }

    // todo: will replace this with __invoke
    public function __get(string $name)
    {
        $name = strtolower($name);
        if ($this->$name) {
            return $this->$name;
        }

        return null;
    }

    // todo: will replace this with __invoke
    public function __set(string $name, $value)
    {
        $this->$name = $value;
    }

    // todo: will remove this
    public function __isset(string $name)
    {
        return isset($this->$name);
    }
}

$array = [
    'hoge' => 1,
    'fuga' => 2,
];

$hoge = new class($array)
{
    use SetterGetterTrait;
};

var_dump($hoge->hoge);
var_dump($hoge->fuga);
<?php

interface HelperInterface
{
}

class HelperTraitException extends \Exception
{
}

class HogeAndFuga implements HelperInterface
{
    public function __invoke(array $hoge = ['hoge'])
    {
        var_dump($hoge);
    }
}

/**
 * Trait HelperTrait
 * @method hoge_and_fuga(string $string)
 */
trait HelperTrait
{
    /**
     * @param string $name
     * @param array $arguments
     * @return mixed
     * @throws HelperTraitException
     */
    public function __call(string $name, array $arguments)
    {
        return $this->call($name, ...$arguments);
    }

    /**
     * @param string $name
     * @param array ...$arguments
     * @return mixed
     * @throws HelperTraitException
     */
    public function call(string $name, ...$arguments)
    {
        $strings = mb_split('_', $name);
        $class_name = '';
        foreach ($strings as $s) {
            $class_name .= ucfirst($s);
        }
        $absolute_path = __NAMESPACE__ . "\\" . $class_name;

        if (($c = new $absolute_path) instanceof HelperInterface) {
            return (new $c)(...$arguments);
        }

        throw new HelperTraitException(sprintf('called undefined class: %s', $c));
    }
}

class Test
{
    use HelperTrait;

    public function main(): void
    {
        $this->hoge_and_fuga('aaa');
    }
}


(new class() { use HelperTrait; })->hoge_and_fuga('bbb');

<?php

require_once __DIR__ . '/../vendor/autoload.php';

use Foo\Bar\SearchRequest;

$foo = new SearchRequest();
$foo->setQuery('hogehoge fugafuga');
$foo->setPageNumber(100);

echo $foo->serializeToString();

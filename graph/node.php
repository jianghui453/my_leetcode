<?php

class Node {
    public $val;
    public $neighbors;
/**
@param Integer $val
@param list<Node> $neighbors*/
    function __construct($val, $neighbors) {
        $this->val = $val;
        $this->neighbors = $neighbors;
    }
}
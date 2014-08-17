<?php

function mySqrt($x) {
	$z = 1.0;
	for ($i = 0; $i < 1000; $i++) {
		$z -= ($z*$z - $x) / (2 * $z);
	}
	return $z;
}

$start = microtime(true);
$loops = 999999;
for ($i = 0; $i < $loops; $i++) {
	mySqrt($i);
}
$end = microtime(true);
print(sprintf("%d loops in %s\n", $loops, $end-$start));
print("======\n");
print(sprintf("%dkb/%dkb\n", memory_get_usage(true)/1024, memory_get_peak_usage(true)/1024));
print("======\n");


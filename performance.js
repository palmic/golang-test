function Sqrt(x) {
        var z = 1.0;
        for (var i = 0; i < 1000; i++) {
                z -= (z*z - x) / (2 * z);
        }   
        return z;
}

var loops = 999999;

var start = process.hrtime();
var memStats = process.memoryUsage();

for (i = 0; i < loops; i++) {
	Sqrt(i);
}   

duration = process.hrtime(start);

console.log(loops + " loops in " + (duration[0] + duration[1]/1000000000).toFixed(6) + "s");
console.log("======");
console.log((memStats.heapUsed/1024) + 'kb/' + memStats.heapTotal/1024 + "kb");
console.log("======");

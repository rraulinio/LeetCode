/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
    var y = x.toString();
    var z = "";
    var neg = false;
    if(-1 * x > 0)
        neg = true;
    for(var i = y.length-1; i >= 0; i--)
       z += y[i];
    var a = parseInt(z);
    if(neg == true)
        a *= -1;
    if(a < -1*Math.pow(2,31) || a > Math.pow(2,31)-1)
        return 0;
    return a;
};

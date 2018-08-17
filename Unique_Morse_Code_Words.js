/**
 * @param {string[]} words
 * @return {number}
 */
function uniqueMorseRepresentations(words) {
	let out = [];
	var code = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."];
    var lts = "abcdefghijklmnopqrstuvwxyz";
    for(v of words){
        var x = v.split("").reduce((total, letter) => total + code[lts.indexOf(letter)], "");
        if(out.indexOf(x) === -1)
            out.push(x);
    }
    return out.length;
};
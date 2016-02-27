
var module = (function() {
	"use strict";
	var fn = function() {
		console.log("test");
	};

	return {
		fn: fn
	};

})();

module.fn();

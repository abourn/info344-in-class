'use strict'

var btn = document.querySelector("button");
var input = document.querySelector("input");
var title = document.querySelector("h1");

btn.addEventListener("click", function() {
	var txt = input.value;
	var url = "http://127.0.0.1:4000/hello" + "?name=" + txt;
	fetch(url)
	.then(function(resp) {
		console.log(resp);
		return resp.text();
	}).then(function(text) {
		title.textContent=text;
	})
}) 
'use strict'

var btn = document.querySelector("button");
var input = document.querySelector("input");
var title = document.querySelector("h1");
var p = document.querySelector("p");
var zipSelect = document.querySelector(".zipSelect");
var stateSelect = document.querySelector(".stateSelect");


btn.addEventListener("click", function() {
	var txt = input.value;
	var url = "http://127.0.0.1:4000/zips/" + txt
	fetch(url)
	.then(function(resp) {
		console.log(resp);
		return resp.json();
	}).then(function(json) {
		var states = [];
		json.forEach(function(element) {
			var zip = element.Code;
			var state = element.State;
			var zipOption = document.createElement('option');

			zipOption.setAttribute('value', zip);
			zipOption.append(zip);
			zipSelect.appendChild(zipOption);

			if (!states.includes(state)) {
				var stateOption = document.createElement('option');

				stateOption.setAttribute('value', state);
				stateOption.append(state);
				stateSelect.appendChild(stateOption);
			}
			states.push(state);

		})

	})
}) 
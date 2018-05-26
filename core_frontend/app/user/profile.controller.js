app.controller('profileController', function profileController($routeParams, userService) {
	var vm = this;

	var username = $routeParams.username;

	vm.getProfile = getProfile;

	function getProfile() {
		userService.getUser(username)
		.then(function(obj) {
			vm.user = obj.data;
		})
		.catch(function(obj) {
			console.log(obj);
		});
	}
	
	function init() {
		vm.getProfile();
	}

	init();
});
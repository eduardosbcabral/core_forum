app.controller('navigationController', function navigationController($scope, userService) {
	var vm = this;

	vm.getLoggedUser = getLoggedUser;

	function getLoggedUser() {
		userService.getLoggedUser('theduardds')
		.then(function(obj) {
			$scope.user = obj.data;
		})
		.catch(function(obj) {
			console.log(obj);
		});
	}
	
	function init() {
		vm.getLoggedUser();
	}

	init();
});
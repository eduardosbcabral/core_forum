app.controller('navigationController', function navigationController($scope, userService, toastr) {
	var vm = this;

	vm.getLoggedUser = getLoggedUser;

	function getLoggedUser() {
		userService.getUser('theduardds')
		.then(function(obj) {
			$scope.user = obj.data;
		})
		.catch(function(obj) {
			toastr.error(obj.data.message, 'Error');
		});
	}
	
	function init() {
		vm.getLoggedUser();
	}

	init();
});
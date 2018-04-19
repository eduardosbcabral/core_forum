app.controller('navigationController', function userController($scope, $location, $timeout) {
	var vm = this;
	
	//vm.getLoggedUser = getLoggedUser;

	$timeout(init, 0);

	/*function getLoggedUser() {
		usuarioService.getLoggedUser('theduardds')
		.then(function(obj) {
			console.log(obj);
			vm.user = obj.data;
		})
		.catch(function(obj) {
			console.log(obj);
		});
	}*/
	
	function init() {
		$('.dropdown-button').dropdown();
	}


	/*function main() {
		vm.getLoggedUser();
	}*/

	//main();
});
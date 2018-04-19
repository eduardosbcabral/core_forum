app.controller('userController', function userController(userService) {
	var vm = this;

	vm.getUsers = getUsers;
	vm.getGenders = getGenders;


	function getUsers() {
		userService.getUsers()
			.then(function(obj) {
				vm.users = obj.data;
			})
			.catch(function(obj){
				console.log(obj);
			});
	}

	function getGenders() {
		userService.getGenders()
			.then(function(obj) {
				vm.genders = obj.data;
			})
			.catch(function(obj){
				console.log(obj);
			});
	}

	function init() {
		vm.getUsers();
		vm.getGenders();
		vm.currentTab = 'User';
		$('ul.tabs').tabs();
	}

	init();
});
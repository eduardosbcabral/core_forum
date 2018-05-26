app.controller('userController', function userController(userService, toastr) {
	var vm = this;

	vm.getUsers = getUsers;
	vm.getGenders = getGenders;
	vm.openModalAddUser = openModalAddUser;
	vm.openModalEditUser = openModalEditUser;
	vm.createUser = createUser;
	vm.editUser = editUser;
	vm.getUser = getUser;
	vm.deleteUser = deleteUser;
	vm.openModalAddGender = openModalAddGender;
	vm.openModalEditGender = openModalEditGender;
	vm.createGender = createGender;
	vm.editGender = editGender;
	vm.getGender = getGender;
	vm.deleteGender = deleteGender;
	vm.limparEscopo = limparEscopo;

	function getUsers() {
		userService.getUsers()
			.then(function(obj) {
				vm.users = obj.data;
			})
			.catch(function(obj){
				console.log(obj);
			});
	}

	function getUser(username) {
		userService.getUser(username)
		.then(function(obj) {
			vm.editedUser = obj.data;
			$(document).ready(function() { 
				Materialize.updateTextFields(); 
				$('select').material_select();

			});

		})
		.catch(function(obj) {
			toastr.error(obj.data.message, 'Error');
		})
	}

	function getGenders() {
		userService.getGenders()
			.then(function(obj) {
				vm.genders = obj.data;
				console.log(vm.genders);
			})
			.catch(function(obj){
				toastr.error(obj.data.message, 'Error');
			});
	}

	function getGender(genderId) {
		userService.getGender(genderId)
		.then(function(obj) {
			vm.editedGender = obj.data;
			$(document).ready(function() { 
				Materialize.updateTextFields(); 
			});

		})
		.catch(function(obj) {
			toastr.error(obj.data.message, 'Error');
		})
	}

	function createUser(user) {
		userService.createUser(user)
		.then(function(obj) {
			vm.getUsers();
			$('#modalAddUser').modal('close');
			vm.limparEscopo();
		})
		.catch(function(obj) {
			toastr.error(obj.data.message, 'Error');
		})
	}

	function editUser(user) {
		userService.editUser(user, vm.username)
		.then(function(obj) {
			vm.getUsers();
			$('#modalEditUser').modal('close');
			vm.limparEscopo();
		})
		.catch(function(obj) {
			toastr.error(obj.data.message, 'Error');
		})
	}

	function deleteUser(user) {
		console.log(user);	
		user.active = false;
		userService.deleteUser(user, user.username)
		.then(function(obj) {
			vm.getUsers();
			vm.limparEscopo();
		})
		.catch(function(obj) {
			toastr.error(obj.data.message, 'Error');
		})
	}

	function createGender(gender) {
		userService.createGender(gender)
		.then(function(obj) {
			vm.getGenders();
			$('#modalAddGender').modal('close');
			vm.limparEscopo();
		})
		.catch(function(obj) {
			toastr.error(obj.data.message, 'Error');
		})
	}

	function editGender(gender) {
		userService.editGender(gender, vm.genderId)
		.then(function(obj) {
			vm.getGenders();
			$('#modalEditGender').modal('close');
			vm.limparEscopo();
		})
		.catch(function(obj) {
			toastr.error(obj.data.message, 'Error');
		})
	}

	function deleteGender(gender) {
		gender.active = false;
		userService.deleteGender(gender, gender._id)
		.then(function(obj) {
			vm.getGenders();
			vm.limparEscopo();
		})
		.catch(function(obj) {
			toastr.error(obj.data.message, 'Error');
		})
	}

	function openModalEditUser(user) {
		vm.limparEscopo();
		$('#modalEditUser').modal();
		$('#modalEditUser').modal('open');
		vm.username = user.username;
		vm.getUser(user.username);
	}

	function openModalAddUser() {
		vm.limparEscopo();
		$('#modalAddUser').modal();
		$('#modalAddUser').modal('open');
	}

	function openModalEditGender(gender) {
		vm.limparEscopo();
		$('#modalEditGender').modal();
		$('#modalEditGender').modal('open');
		vm.genderId = gender._id;
		vm.getGender(vm.genderId);
	}

	function openModalAddGender() {
		vm.limparEscopo();
		$('#modalAddGender').modal();
		$('#modalAddGender').modal('open');
	}

	function limparEscopo() {
		delete vm.user	
		delete vm.editedUser
		
		delete vm.gender
		delete vm.editedGender
	}

	function init() {
		vm.getUsers();
		vm.getGenders();
		vm.currentTab = 'User';
		$('ul.tabs').tabs();
	}

	init();

	angular.element(document).ready(function () {
  		$('select').material_select();
	});
});
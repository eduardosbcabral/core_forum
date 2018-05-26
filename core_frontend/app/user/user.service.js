app.service('userService', function userService($http, configService) {

	return {
		getUsers: getUsers,
		getGenders: getGenders,
		getUser: getUser,
		createUser: createUser,
		editUser: editUser,
		deleteUser: deleteUser,
		getGender: getGender,
		createGender: createGender,
		editGender: editGender,
		deleteGender: deleteGender
	};

	function getUsers() {
		return $http.get(configService.urlApi + 'user');
	}
	
	function getGenders() {
		return $http.get(configService.urlApi + 'gender');
	}

	function getUser(username) {
		var request = {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
            url: configService.urlApi + 'user/' + username,
        }
        return $http(request);
	}

	function getGender(genderId) {
		var request = {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
            url: configService.urlApi + 'gender/' + genderId,
        }
        return $http(request);
	}

	function createUser(user) {
		var request = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            data: user,
            url: configService.urlApi + 'user',
        }
        return $http(request);
	}

	function editUser(user, username) {
		var request = {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            data: user,
            url: configService.urlApi + 'user/' + username,
        }
        return $http(request);
	}

	function deleteUser(user, username) {
		var request = {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            data: user,
            url: configService.urlApi + 'user/' + username,
        }
        return $http(request);
	}

	function createGender(gender) {
		var request = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            data: gender,
            url: configService.urlApi + 'gender',
        }
        return $http(request);
	}

	function editGender(gender, genderId) {
		var request = {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            data: gender,
            url: configService.urlApi + 'gender/' + genderId,
        }
        return $http(request);
	}

	function deleteGender(gender, genderId) {
		var request = {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            data: gender,
            url: configService.urlApi + 'gender/' + genderId,
        }
        return $http(request);
	}
});
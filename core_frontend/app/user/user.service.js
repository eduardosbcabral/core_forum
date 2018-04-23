app.service('userService', function userService($http, configService) {

	return {
		getUsers: getUsers,
		getGenders: getGenders,
		getLoggedUser: getLoggedUser
	};

	function getUsers() {
		return $http.get(configService.urlApi + 'user');
	}
	
	function getGenders() {
		return $http.get(configService.urlApi + 'gender');
	}

	function getLoggedUser(username) {
		var request = {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
            url: configService.urlApi + 'user/' + username,
        }
        return $http(request);
	}
});
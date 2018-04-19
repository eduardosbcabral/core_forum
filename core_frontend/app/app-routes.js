app.config(function config($routeProvider, $locationProvider) {

	$routeProvider
		.when('/user', {
			templateUrl: '/app/user/user.html',
			controller: 'userController',
			controllerAs: 'vm',
		});

	$routeProvider.otherwise({redirectTo: '/'});

	$locationProvider.hashPrefix('');
	}
);
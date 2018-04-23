app.config(function config($routeProvider, $locationProvider) {

	$routeProvider
		.when('/user', {
			templateUrl: '/app/user/user.html',
			controller: 'userController',
			controllerAs: 'vm',
		})
		.when('/user/:username', {
			templateUrl: '/app/user/profile.html',
			controller: 'profileController',
			controllerAs: 'vm',
		})
		.when('/category', {
			templateUrl: '/app/category/category.html',
			controller: 'categoryController',
			controllerAs: 'vm',
		})
		.when('/category/:categoryId/thread', {
			templateUrl: '/app/thread/thread.html',
			controller: 'threadController',
			controllerAs: 'vm',
		})

	$routeProvider.otherwise({redirectTo: '/'});

	$locationProvider.hashPrefix('');
	}
);
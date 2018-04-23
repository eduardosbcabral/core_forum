app.service('categoryService', function categoryService($http, configService) {

	return {
		getCategories: getCategories
	};

	function getCategories() {
		return $http.get(configService.urlApi + 'category');
	}
});
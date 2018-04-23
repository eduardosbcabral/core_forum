app.service('threadService', function threadService($http, configService) {

	return {
		getThreads: getThreads
	};

	function getThreads(categoryId) {
		return $http.get(configService.urlApi + categoryId + '/thread');
	}
});
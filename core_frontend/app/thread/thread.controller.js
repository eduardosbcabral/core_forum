app.controller('threadController', function threadController($routeParams, threadService) {
	var vm = this;

	var categoryId = $routeParams.categoryId;

	vm.getThreads = getThreads;

	function getThreads() {
		threadService.getThreads(categoryId)
			.then(function(obj) {
				vm.threads = obj.data;
			})
			.catch(function(obj){
				console.log(obj);
			});
	}

	function init() {
		vm.getThreads();
	}

	init();
});
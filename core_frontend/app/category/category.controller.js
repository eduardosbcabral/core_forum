app.controller('categoryController', function categoryController(categoryService) {
	var vm = this;

	vm.getCategories = getCategories;

	function getCategories() {
		categoryService.getCategories()
			.then(function(obj) {
				vm.categories = obj.data;
			})
			.catch(function(obj){
				console.log(obj);
			});
	}

	function init() {
		vm.getCategories();
	}

	init();
});
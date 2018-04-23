/*================================================================================
  Item Name: Materialize - Material Design Admin Template
  Version: 4.0
  Author: PIXINVENT
  Author URL: https://themeforest.net/user/pixinvent/portfolio
================================================================================*/

  /*Preloader*/
  $(window).on('load', function() {
    setTimeout(function() {
      $('body').addClass('loaded');
    }, 200);
  });

  $(function() {

    "use strict";

    var window_width = $(window).width();
    var openIndex;

    // Search class for focus
    $('.header-search-input').focus(
      function() {
        $(this).parent('div').addClass('header-search-wrapper-focus');
      }).blur(
      function() {
        $(this).parent('div').removeClass('header-search-wrapper-focus');
      });

    // Check first if any of the task is checked
    $('#task-card input:checkbox').each(function() {
      checkbox_check(this);
    });

    // Task check box
    $('#task-card input:checkbox').change(function() {
      checkbox_check(this);
    });

    // Check Uncheck function
    function checkbox_check(el) {
      if (!$(el).is(':checked')) {
        $(el).next().css('text-decoration', 'none'); // or addClass
      } else {
        $(el).next().css('text-decoration', 'line-through'); //or addClass
      }
    }

    // Swipeable Tabs Demo Init
    if ($('#tabs-swipe-demo').length) {
      $('#tabs-swipe-demo').tabs({
        'swipeable': true
      });
    }

    // Plugin initialization

    $('select').material_select();
    // Set checkbox on forms.html to indeterminate
    var indeterminateCheckbox = document.getElementById('indeterminate-checkbox');
    if (indeterminateCheckbox !== null)
      indeterminateCheckbox.indeterminate = true;

    // Materialize Slider
    $('.slider').slider({
      full_width: true
    });

    // Commom, Translation & Horizontal Dropdown
    $('.dropdown-button, .dropdown-menu').dropdown({
      inDuration: 300,
      outDuration: 225,
      constrainWidth: false,
      hover: true,
      gutter: 0,
      belowOrigin: true,
      alignment: 'left',
      stopPropagation: false
    });
    // Notification, Profile & Settings Dropdown
    $('.profile-button, .dropdown-settings').dropdown({
      inDuration: 300,
      outDuration: 225,
      constrainWidth: false,
      hover: true,
      gutter: 0,
      belowOrigin: true,
      alignment: 'right',
      stopPropagation: false
    });

    // Materialize Tabs
    $('.tab-demo').show().tabs();
    $('.tab-demo-active').show().tabs();

    // Materialize Parallax
    $('.parallax').parallax();

    // Materialize scrollSpy
    $('.scrollspy').scrollSpy();

    // Materialize tooltip
    $('.tooltipped').tooltip({
      delay: 50
    });

    //Main Left Sidebar Menu
    $('.sidebar-collapse').sideNav({
      edge: 'left', // Choose the horizontal origin
    });

    // Overlay Menu (Full screen menu)
    $('.menu-sidebar-collapse').sideNav({
      menuWidth: 240,
      edge: 'left', // Choose the horizontal origin
      //closeOnClick:true, // Set if default menu open is true
      menuOut: false // Set if default menu open is true
    });

    //Main Left Sidebar Chat
    $('.chat-collapse').sideNav({
      menuWidth: 300,
      edge: 'right',
    });

    // Pikadate datepicker
    $('.datepicker').pickadate({
      selectMonths: true, // Creates a dropdown to control month
      selectYears: 15 // Creates a dropdown of 15 years to control year
    });

    // Floating-Fixed table of contents (Materialize pushpin)
    if ($('nav').length) {
      $('.toc-wrapper').pushpin({
        top: $('nav').height()
      });
    } else if ($('#index-banner').length) {
      $('.toc-wrapper').pushpin({
        top: $('#index-banner').height()
      });
    } else {
      $('.toc-wrapper').pushpin({
        top: 0
      });
    }

    // Toggle Flow Text
    var toggleFlowTextButton = $('#flow-toggle')
    toggleFlowTextButton.click(function() {
      $('#flow-text-demo').children('p').each(function() {
        $(this).toggleClass('flow-text');
      })
    });

    //Alerts
    $("#card-alert .close").click(function() {
      $(this).closest('#card-alert').fadeOut('slow');
    });

    //Toggle Containers on page
    var toggleContainersButton = $('#container-toggle-button');
    toggleContainersButton.click(function() {
      $('body .browser-window .container, .had-container').each(function() {
        $(this).toggleClass('had-container');
        $(this).toggleClass('container');
        if ($(this).hasClass('container')) {
          toggleContainersButton.text("Turn off Containers");
        } else {
          toggleContainersButton.text("Turn on Containers");
        }
      });
    });

    // Detect touch screen and enable scrollbar if necessary
    function is_touch_device() {
      try {
        document.createEvent("TouchEvent");
        return true;
      } catch (e) {
        return false;
      }
    }
    if (is_touch_device()) {
      $('#nav-mobile').css({
        overflow: 'auto'
      })
    }
  });
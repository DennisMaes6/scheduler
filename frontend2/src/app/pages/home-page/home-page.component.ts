import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  navLinks: any[];
  activeLinkIndex = -1;

  constructor(private router: Router) {
    this.navLinks = [
      {
        label: 'HomePage',
        link: '/home-page',
        index: 0
      }, {
        label: 'MainPage',
        link: '/main',
        index: 1
      }, {
        label: 'ModelParametersPage',
        link: '/model-parameters',
        index: 2
      }, {
        label: 'InstanceDataPage',
        link: '/instance-data',
        index: 3
      }, {
        label: 'ScheduleViewPage',
        link: '/schedule-view',
        index: 4
      }
    ]
   }

  ngOnInit(): void {
    console.log(this.router.url)
    console.log(this.navLinks.indexOf(this.navLinks.find(tab => tab.link ===  this.router.url)))
    this.router.events.subscribe((res) => {
      this.activeLinkIndex = this.navLinks.indexOf(this.navLinks.find(tab => tab.link === '.' + this.router.url));
    });
  }

}

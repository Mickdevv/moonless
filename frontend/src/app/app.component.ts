import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MenuItem } from 'primeng/api';
import { Menubar } from 'primeng/menubar';
import { ButtonModule } from 'primeng/button';
@Component({
  selector: 'app-root',
  imports: [RouterOutlet, Menubar, ButtonModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {
  title = 'frontend';

  items: MenuItem[] | undefined;

  ngOnInit() {
    this.items = [
      {
        label: 'Home',
        icon: 'pi pi-home',
        routerLink: '/home'

      },
      {
        label: 'Pictures',
        icon: 'pi pi-star',
        routerLink: '/pictures'
      },
      {
        label: 'Songs',
        icon: 'pi pi-star',
        routerLink: '/songs'
      },
      {
        label: 'Videos',
        icon: 'pi pi-star',
        routerLink: '/videos'
      },
      {
        label: 'Events',
        icon: 'pi pi-star',
        routerLink: '/events'
      },

      {
        label: 'Contact',
        icon: 'pi pi-envelope',
        routerLink: '/contact'
      }
    ]
  }
}

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
        url: '/home'

      },
      {
        label: 'Pictures',
        icon: 'pi pi-star',
        url: '/pictures'
      },
      {
        label: 'Songs',
        icon: 'pi pi-star',
        url: '/songs'
      },
      {
        label: 'Videos',
        icon: 'pi pi-star',
        url: '/videos'
      },
      {
        label: 'Events',
        icon: 'pi pi-star',
        url: '/events'
      },

      {
        label: 'Contact',
        icon: 'pi pi-envelope',
        url: '/contact'
      }
    ]
  }
}

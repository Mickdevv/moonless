import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { SongsComponent } from './songs/songs.component';
import { PicturesComponent } from './pictures/pictures.component';
import { VideosComponent } from './videos/videos.component';
import { EventsComponent } from './events/events.component';
import { ContactComponent } from './contact/contact.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';

export const routes: Routes = [
    { path: "home", component: HomeComponent },
    { path: "songs", component: SongsComponent },
    { path: "pictures", component: PicturesComponent },
    { path: "videos", component: VideosComponent },
    { path: "events", component: EventsComponent },
    { path: "contact", component: ContactComponent },
    { path: "**", component: PageNotFoundComponent },
];

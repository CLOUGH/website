import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { HomeRoutingModule } from './home-routing.module';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { CoreModule } from 'src/app/core/core.module';


@NgModule({
  declarations: [
    HomePageComponent
  ],
  imports: [
    CoreModule,
    CommonModule,
    HomeRoutingModule
  ]
})
export class HomeModule { }

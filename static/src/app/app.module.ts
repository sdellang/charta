import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { NamespacesService } from './services/namespaces.service'
import { PropertiesService } from './services/properties.service'
import { DisplayPropertiesService } from './services/displayProperties.service'

import { AppComponent } from './app.component';
import { NamespacesListComponent } from './components/namespace-list.component'
import { PropertyListComponent } from './components/property-list.component'

import { VarsPipe } from './vars.pipe'

@NgModule({
  declarations: [
    AppComponent,
    NamespacesListComponent,
    PropertyListComponent,   
    VarsPipe
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule
  ],
  providers: [
    NamespacesService,
    PropertiesService,
    DisplayPropertiesService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

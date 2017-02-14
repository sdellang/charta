import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { PropertiesService } from '../services/properties.service'
import { DisplayPropertiesService } from '../services/displayProperties.service'
import {Pod} from '../model/pod'

@Component({
  selector: 'properties',
  templateUrl: './property-list.component.html',
  styleUrls: ['../css/main.component.css']
})

export class PropertyListComponent {
    pods: Pod[]

    constructor (private propertiesService: PropertiesService,
                 private displayPropertiesService: DisplayPropertiesService) {}

    ngOnInit(): void {
        this.displayPropertiesService.getProperties()
        .subscribe(retns => this.getNsProperties(retns))
    }

    getNsProperties(ns: string): void {
        console.log("getting properties for: "+ns)
        this.propertiesService.getProperties(ns)
        .subscribe(retPods => this.pods = retPods)
    }

    
}
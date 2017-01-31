import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { PropertiesService } from '../services/properties.service'
import { DisplayPropertiesService } from '../services/displayProperties.service'
import {Pod} from '../model/pod'

@Component({
  selector: 'properties',
  templateUrl: './property-list.component.html'
})

export class PropertyListComponent {
    pods: Pod[]
    visible: boolean

    constructor (private propertiesService: PropertiesService,
                 private displayPropertiesService: DisplayPropertiesService) {}

    ngOnInit(): void {
        this.visible = false
        this.displayPropertiesService.getProperties()
        .subscribe(retns => this.getNsProperties(retns))
    }

    getNsProperties(ns: string): void {
        this.propertiesService.getProperties(ns)
        .subscribe(retPods => this.pods = retPods)
    }

    
}
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { NamespacesService } from '../services/namespaces.service'
import { DisplayPropertiesService } from '../services/displayProperties.service'

@Component({
  selector: 'namespaces-card',
  templateUrl: './card-namespace.component.html',
  styleUrls: ['../css/main.component.css']
})

export class CardNamespaceComponent {
    namespaces: string[]

    constructor (private namespacesService: NamespacesService,
                 private displayPropertiesService: DisplayPropertiesService) {}

    ngOnInit(): void {
        this.namespacesService.getNamespaces()
        .subscribe(retns => this.namespaces = retns)

    }

    loadProperties(namespace: string): void {
        this.displayPropertiesService.setProperties(namespace)
    }
    
}
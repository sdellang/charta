import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { Router } from '@angular/router';

import { NamespacesService } from '../services/namespaces.service'
import { DisplayPropertiesService } from '../services/displayProperties.service'

@Component({
  selector: 'namespaces',
  templateUrl: './namespace-list.component.html'
})

export class NamespacesListComponent {
    namespaces: string[]

    @Output() onPropListVisible = new EventEmitter<boolean>();

    constructor (private namespacesService: NamespacesService,
                 private displayPropertiesService: DisplayPropertiesService) {}

    ngOnInit(): void {
        this.namespacesService.getNamespaces()
        .subscribe(retns => this.namespaces = retns)

    }

    loadProperties(namespace: string): void {
        
        this.displayPropertiesService.setProperties(namespace)
        this.onPropListVisible.emit(true);
        console.log("loadProperties")
    }
    
}
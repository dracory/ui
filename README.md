# UI <a href="https://gitpod.io/#https://github.com/gouniverse/ui" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

[![Tests Status](https://github.com/gouniverse/ui/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/gouniverse/ui/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/ui)](https://goreportcard.com/report/github.com/gouniverse/ui)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/ui)](https://pkg.go.dev/github.com/gouniverse/ui)

## License


This project is licensed under the GNU Affero General Public License v3.0 (AGPL-3.0). You can find a copy of the license at [https://www.gnu.org/licenses/agpl-3.0.en.html](https://www.gnu.org/licenses/agpl-3.0.txt)

For commercial use, please use my [contact page](https://lesichkov.co.uk/contact) to obtain a commercial license.

## Introduction

This package allows to build user interfaces based on blocks.

The block based user interface design approach breaks down complex user interfaces into smaller, self-contained blocks or modules.

These blocks can be arranged and rearranged to create different layouts and user experiences.

## Benefits of Block-Based User Interfaces

- **Faster Development:**
  Pre-built blocks streamline the UI creation process, reducing development time.
  Developers can focus on the logic and functionality rather than the nitty-gritty of UI design.
  
- **Reduced Technical Barrier:**
  No need for in-depth knowledge of complex UI frameworks or languages.
  Visual nature makes it accessible to a wider range of users, including non-programmers.

- **Consistency and Standardization:**
  Enforces consistency in UI design and layout across different parts of the application.
  Promotes reusability of UI components.
  
- Flexibility:
  Can be adapted to various UI paradigms (e.g., web, mobile, desktop) by customizing the available blocks.
  Offers the potential for creating complex and dynamic UIs.

## Disadvantages of Block-Based User Interfaces
Before starting on this path beware of the following:

- **Limited Customization:**
  While blocks provide flexibility, they may not always meet specific design requirements.
  May not be suitable for highly customized or unique UIs.
  
- **Steeper Learning Curve:**
  While easier than traditional UI development, there's still a learning curve to understand the block library and its capabilities.
  
- **Performance Overhead:**
  In some cases, block-based UIs might introduce performance overhead due to the underlying framework and interpretation of blocks.
  
- **Vendor Lock-in:**
  Reliance on a specific block-based framework can limit future options and migration possibilities.

> Overall, block-based user interfaces offer a promising approach to streamline UI development, especially for teams with diverse skill sets. However, their suitability depends on the specific project requirements and the trade-offs between speed, flexibility, and customization.

## Installation
```
go get -u github.com/gouniverse/ui
```

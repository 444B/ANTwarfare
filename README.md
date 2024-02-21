# ANTwarfare

## TL;DR
Phase 1: For a POC, we will apply ANT to foxhole by using the warapi as a data source to understand effective strategies and insights.  
Phase 2: The Long Term Vision is to integrate OpenAI/(insert other LLM) and being able to ask how to achieve stated goals, based on the understanding provided by ANTwarfare.   
Phase 3: The Moonshot Vision is to provide in real time decentralized coordination framework for Officers to receive insights for strategizing and Operatives will receive push notifications to carry out objective missions.  
Right now, we are focusing on refining correct theory and taking baby steps, which we would be glad to have more volunteers involved since we have alot of code scrooping to do  

## Introduction
Welcome to ANTwarfare, an innovative project that applies computational modeling and [Actor-Network Theory (ANT)](https://en.wikipedia.org/wiki/Actor%E2%80%93network_theory) to the complex, multiplayer game environment of [Foxhole](https://foxhole.wiki.gg/wiki/Foxhole_Wiki). This initiative aims to dissect and enhance understanding of the intricate interactions among players, resources, territories, and infrastructures within the game. Through ANTwarfare, we explore effective strategies, predict outcomes of specific actions, and refine gameplay strategies, all while contributing to the broader discourse on network theory in complex systems.  

> [!NOTE]   
This is an early stage project and the current code is still based in theory. There will be an update in the projects [Discussions](https://github.com/444B/ANTwarfare/discussions) once we have a Proof of Concept (POC) that achieves one / many of the [objectives](https://github.com/444B/ANTwarfare?tab=readme-ov-file#objectives), listed below.  
> Please feel free to get involved or share the project, which is intentionally set up for open source collaboration and a community effort.  
> Make sure to take a look at our [Roadmap](https://github.com/users/444B/projects/4/views/4)  
> The current main issue for developing the POC is tracked [here](https://github.com/444B/ANTwarfare/issues/8)   

## Background
Foxhole is a multiplayer game that requires strategic collaboration and competition, involving dynamic interactions that make it an ideal candidate for computational modeling and ANT analysis.  
ANT, a framework that treats humans and non-humans (e.g., technologies, objects) equally as actors in a network, offers a unique lens through which to examine the game's social and material dynamics.  

## Objectives  
- **To identify effective strategies** for players within the Foxhole game environment.  
- **To predict the outcomes** of specific player actions and interactions within the game.  
- **To refine the model** based on actual gameplay outcomes, enhancing the predictive accuracy and strategic insights provided by the project.  
- **To explore the implications** of ANT in understanding and strategizing within complex, multiplayer game systems.  

## Getting Involved  
There are a few ways to get involved, which are detailed in the [CONTRIBUTING.md](/.github/CONTRIBUTING.md)   
Briefly, they are:  
1. Writing code, squashing bugs and developing features  
1. Sharing your foxhole experience and helping us map it to ANT by engaging in [Discussions](https://github.com/444B/ANTwarfare/discussions)   
1. Writing docs in the [Wiki](https://github.com/444B/ANTwarfare/wiki)  
1. (once POC is functional) - Testing the [Security](https://github.com/444B/ANTwarfare/security) of the project / associated code and reporting

## Contribution Guidelines  
We welcome contributions to the ANTwarfare project! Whether you're looking to fix bugs, enhance the simulation engine, or offer strategic insights, here's how you can contribute:  

- **Fork the repository** and create your feature branch: `git checkout -b my-new-feature`  
- **Commit your changes**: `git commit -am 'Add some feature'`  
- **Push to the branch**: `git push origin my-new-feature`  
- **Submit a pull request**: Ensure your contributions are well-documented and include any relevant test results.  

For more detailed information, please refer to our [CONTRIBUTING.md](.github/CONTRIBUTING.md).  

## Project Design
### Language
Go's concurrency features, along with structs and interfaces, can effectively model real-time dynamics and is suitable for the application of ANT to Foxhole:  
- Structs for Actors: Define structs for each actor type (players, resources, territories) with fields representing their state and properties.  
- Interfaces for Actions: Use interfaces to define common actions (e.g., attack, defend, gather) that actors can perform, allowing for polymorphism.  
- Concurrency for Real-Time Dynamics: Utilize Go's goroutines and channels to simulate concurrent actions and interactions among actors, reflecting the game's real-time nature.
  
### Data
The main way of understanding the Actors is from data collected via the Official [Foxhole WarAPI v1](https://github.com/clapfoot/warapi) however additional data sources should be found and integrated for a full scope of ANT as applied to Foxhole (Meme warfare on [/r/foxholegame](https://www.reddit.com/r/foxholegame/), Discord, etc)  

### Theory

#### Strengths of ANT for This Project

- **Holistic View**: ANT's strength lies in its ability to provide a comprehensive view of a network, considering both human (players, factions) and non-human (territories, resources) actors equally. This can be particularly powerful in analyzing a game like Foxhole, where success depends on a complex interplay of strategy, geography, resources, and player collaboration.  

- **Flexibility**: ANT does not prescribe a specific method for analysis but rather offers a lens through which to view the data. This flexibility can be advantageous when dealing with the diverse and dynamic data provided by the War API.  

- **Emphasis on Relationships**: ANT's focus on relationships and interactions can reveal insights into how different elements of the game influence each other, offering a deeper understanding of game dynamics beyond mere statistical analysis.  

#### Challenges and Considerations  

- **Complexity**: ANT can be complex to implement, requiring a detailed mapping and analysis of actors and their interactions. This might necessitate a significant investment in data collection, processing, and interpretation.  

- **Data Limitations**: The effectiveness of ANT analysis is contingent on the availability and granularity of data. While the War API provides a wealth of information, there may be limitations on the data necessary to fully explore certain aspects of the game through an ANT lens.  

- **Interpretation and Application**: Insights derived from ANT analysis might be abstract or require nuanced interpretation to translate into actionable strategies or understandings. Ensuring these insights are meaningful and applicable to your goals or the interests of the Foxhole player community might require additional effort.    

## Contribution Guidelines  
We welcome contributions to the ANTwarfare project! Whether you're looking to fix bugs, enhance the simulation engine, or offer strategic insights, here's how you can contribute:  

- **Fork the repository** and create your feature branch: `git checkout -b my-new-feature`  
- **Commit your changes**: `git commit -am 'Add some feature'`  
- **Push to the branch**: `git push origin my-new-feature`  
- **Submit a pull request**: Ensure your contributions are well-documented and include any relevant test results.  

For more detailed information, please refer to our [CONTRIBUTING.md](.github/CONTRIBUTING.md).  

## License  
ANTwarfare is released under the GNU General Public License v3.0 (GPL-3.0), ensuring that its contributions remain open-source and freely available. For more details, see the [LICENSE](.github/LICENSE) file.  

## Acknowledgments  
- Thank you to all the contributors who have shared their insights and expertise.  
- Special thanks to the Foxhole community for inspiring this project.  

![ANTWarfare DALLE art](/markdown_assets/background.png)  

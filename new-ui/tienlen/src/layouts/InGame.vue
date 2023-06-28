<template>
  <div>
    <div class="gameTable">
      <slot name="gameTable" />
    </div>
    <div class="headsUp">
      <div class="scoreArea">
        <slot name="scoreArea" />
      </div>
      <div class="controlsArea">
        <slot name="controlsArea" />
      </div>
      <div class="activityArea">
        <slot name="activityArea" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.gameTable {
  height: 750px;
  width: 1500px;
  padding: 10px;
  background-color: rgba(48, 112, 16, 0.7);
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  border: 4mm ridge rgba(170, 50, 50, 0.6);
  border-radius: 5%;
}

.headsUp {
  width: 1540px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin: 15px auto 0 auto;

  & .scoreArea {
    width: 510px;
    height: 200px;
    background-color: rgba(255, 255, 255, 0.6);
    border: 3px black double;
  }

  & .controlsArea {
    width: 170px;
    height: 200px;
    background-color: rgba(50, 50, 50, 0.8);
    border: 3px black double;
  }

  & .activityArea {
    width: 800px;
    height: 200px;
    background-color: rgba(50, 50, 50, 0.8);
    border: 3px black double;
  }
}
</style>

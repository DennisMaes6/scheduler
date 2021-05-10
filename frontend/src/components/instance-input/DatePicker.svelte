<script lang=typescript>
  import Modal from "../Modal.svelte"
  import Button from "../model-input/Button.svelte"
import DayHeader from "../schedule-view/DayHeader.svelte";

  export let currentDate: {day: number, month: number, year: number};
  export let handleSubmit: VoidFunction;
  export let setStartDate: (date: Date) => void;

  let visibleMonth = currentDate.month
  let visibleYear = currentDate.year
  let blankDays: Date[];
  let monthDays: Date[];

  const MONTH_NAMES = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
  const DAYS = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];

  function updateMonthDays() {
    blankDays = Array((new Date(visibleYear, visibleMonth-1).getDay()+6) % 7);
    monthDays = new Array(31)
      .fill('')
      .map((_,i) => new Date(visibleYear,visibleMonth-1,i+1))
      .filter((v) => v.getMonth() === visibleMonth-1)
  }

  function incrementMonth() {
    if (visibleMonth == 12) {
      visibleMonth = 1
      visibleYear++
    } else {
      visibleMonth++
    }
    updateMonthDays()
  }
    

  function decrementMonth() {
    if (visibleMonth == 1) {
      visibleMonth = 12;
      visibleYear--
    } else {
      visibleMonth--
    }
    updateMonthDays()
  }

  updateMonthDays()

  function resetStateAnd(callback: VoidFunction) {
    visibleMonth = currentDate.month
    visibleYear = currentDate.year
    updateMonthDays()
    callback()
  }

</script>


<Modal>
  <div class="relative mt-4" slot="trigger" let:open  on:click={() => resetStateAnd(open)}> 
    <input 
    class="p-2 w-full leading-none rounded-lg shadow-sm focus:outline-none focus:shadow-outline text-gray-600 font-medium text-xs"
    type="text"
    value={new Date(currentDate.year, currentDate.month-1, currentDate.day).toDateString()}
    readonly
    placeholder="Select date">

    <div class="absolute top-0 right-0 p-1.5">
      <svg class="h-5 w-5 text-gray-400"  fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"/>
      </svg>
    </div>
  </div>

  <div class="my-2" slot="header">
      <h1 class="font-bold"> Select start date </h1>
  </div>

  <div slot="content" class="w-64 h-72 bg-white m-6 rounded-lg border border-black p-4">
    <div class="flex justify-between items-center">
      <div>
        <span class="text-lg font-bold text-gray-800"> {MONTH_NAMES[visibleMonth-1]}</span>
        <span class="ml-1 text-lg text-gray-600 font-normal"> {visibleYear}</span>
      </div>
      <div>
        <div 
          type="button"
          class="inline-flex cursor-pointer hover:bg-gray-200 p-1 rounded-full" 
          on:click={() => decrementMonth()}>
          <svg class="h-6 w-6 text-gray-500 inline-flex"  fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
          </svg>  
        </div>
        <div 
          type="button"
          class="inline-flex cursor-pointer hover:bg-gray-200 p-1 rounded-full" 
          on:click={() => incrementMonth()}>
          <svg class="h-6 w-6 text-gray-500 inline-flex"  fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
          </svg>									  
        </div>
      </div>
    </div>

    <div class="flex flex-wrap my-3 -mx-1">
      {#each DAYS as day}
        <div style="width: 14.26%" class="px-1">
          <div class="text-gray-800 font-medium text-center text-xs"> {day} </div>
        </div>
      {/each}
    </div>

    <div class="flex flex-wrap -mx-1">
      {#each blankDays as _}
        <div style="width: 14.28%"/>
      {/each}
      {#each monthDays as day}
        <div style="width: 14.28%" class="px-0.5 mb-1">
          {#if day.getDay() == 5}
            <div 
              class={"cursor-pointer leading-loose text-center text-sm rounded-full font-bold " + 
                      (currentDate.year === day.getFullYear()
                        && currentDate.month-1 === day.getMonth()
                        && currentDate.day === day.getDate() 
                      ? "bg-green-500 text-white" 
                      : "text-gray-700 hover:bg-green-300")
                    }
              on:click={() => setStartDate(day)}
            >
              {day.getDate()}
            </div>   
          {:else}
            <div class={"leading-loose text-center text-sm rounded-full text-gray-500 font-normal"}>     
              {day.getDate()}
            </div>
          {/if}
        </div>
      {/each}
    </div>
  </div>

  <div class="my-5 flex flex-row justify-end space-x-2" slot="footer" let:close>
      <Button primary={false} callback={close}> Close </Button>
      <Button callback={() => {handleSubmit(); close()}}> Submit </Button>
  </div>
</Modal>
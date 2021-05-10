<script lang="typescript">
    import { get_spread_object, get_spread_update } from "svelte/internal"
import type { Day } from "../../openapi"

    import Modal from "../Modal.svelte"
    import Button from "../model-input/Button.svelte"

    export let days: Day[]
    let months: {month: number, year: number}[] = Array.from(
        new Set(days.map((d) => {return {month: d.date.month, year: d.date.year}}))
    ).sort((a, b) => {
        if (a.year == b.year) {
            return a.month - b.month
        }
        return a.year - b.year
    })
    
    export let handleSubmit: VoidFunction


    function toggleHoliday(day: Day) {
        day.is_holiday = !day.is_holiday
        days = days
        updateMonthDays()
    }

    let visibleMonth = days[0].date.month
    let visibleYear = days[0].date.year
    let blankDays: number[];
    let monthDays: Date[];

    const MONTH_NAMES = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
    const DAYS = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];

    function updateMonthDays() {
        blankDays = Array((new Date(visibleYear, visibleMonth-1).getDay()+6) % 7);
        monthDays = new Array(31)
            .fill('')
            .map((_,i) => new Date(visibleYear,visibleMonth-1,i+1))
            .filter((v) => v.getMonth() === visibleMonth-1);
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
        visibleMonth = days[0].date.month
        visibleYear = days[0].date.year
        updateMonthDays()
        callback()
    }

    function isPartOfPlanningPeriod(date: Date): boolean {
        return days.some(d => d.date.day === date.getDate() && d.date.month === date.getMonth() + 1 && d.date.year === date.getFullYear())
    }

    function getDay(date: Date): Day {
        return days.find(d => d.date.day === date.getDate() && d.date.month === date.getMonth() + 1 && d.date.year === date.getFullYear())
    }
</script>



<Modal>
    <div class="mt-4" slot="trigger" let:open> 
        <Button callback={() => resetStateAnd(open)}> Edit holidays </Button>
    </div>
    <div class="my-2" slot="header">
        <h1 class="font-bold"> Edit holidays </h1>
    </div>
    <div slot="content" class="w-64 h-72 bg-white m-6 rounded-lg border border-black p-4">
        <div class="flex justify-between items-center">
            <div>
                <span class="text-lg font-bold text-gray-800"> {MONTH_NAMES[visibleMonth-1]}</span>
                <span class="ml-1 text-lg text-gray-600 font-normal"> {visibleYear}</span>
            </div>
            <div class="flex flex-row">
                {#if visibleYear === months[0].year && visibleMonth === months[0].month}
                    <div class="h-8 w-8"></div>
                {:else}
                    <div 
                        type="button"
                        class="inline-flex cursor-pointer hover:bg-gray-200 p-1 rounded-full" 
                        on:click={() => decrementMonth()}>
                        <svg class="h-6 w-6 text-gray-500 inline-flex"  fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
                        </svg>  
                    </div>
                {/if}
                {#if visibleYear === months[months.length-1].year && visibleMonth === months[months.length-1].month}
                    <div class="h-8 w-8"></div>
                {:else}
                    <div 
                        type="button"
                        class="inline-flex cursor-pointer hover:bg-gray-200 p-1 rounded-full" 
                        on:click={() => incrementMonth()}>
                        <svg class="h-6 w-6 text-gray-500 inline-flex"  fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
                        </svg>									  
                    </div>
                {/if}
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
          {#each monthDays as date}
            {#if isPartOfPlanningPeriod(date)}
                <div style="width: 14.28%" class="px-0.5 mb-1">
                    <div 
                        class={"cursor-pointer leading-loose text-center font-bold text-sm rounded-full text-gray-700 " + 
                                (getDay(date).is_holiday
                                ? "bg-yellow-200" 
                                : "hover:bg-yellow-100")
                                }
                        on:click={() => toggleHoliday(getDay(date))}
                        >
                        {getDay(date).date.day}
                    </div>   
                </div>
            {:else}
                <div style="width: 14.28%" class="px-0.5 mb-1">
                    <div class="leading-loose text-center text-sm rounded-full text-gray-700">
                        {date.getDate()}
                    </div>   
                </div>
            {/if}
          {/each}
        </div>
      </div>
    <div class="my-5 flex flex-row justify-end space-x-2" slot="footer" let:close>
        <Button primary={false} callback={close}> Close </Button>
        <Button callback={() => {handleSubmit(); close()}}> Submit </Button>
    </div>
</Modal>
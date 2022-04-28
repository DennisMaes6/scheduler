<script lang=typescript>

    import type { Assignment } from '../../openapi'
    import { ShiftType } from '../../openapi'

    export let assignment: Assignment
    export let free_day: boolean

    function getColor(st: ShiftType): string {
        switch (st) {
            case ShiftType.JAEV:
                return "bg-yellow-400";
            case ShiftType.SANW:
                return "bg-green-400";
            case ShiftType.JAWE:
                return "bg-red-500";
            case ShiftType.JAHO:
                return "bg-blue-500";
            case ShiftType.SAEW:
                return "bg-green-500";
            case ShiftType.SAWE:
                return "bg-red-600";
            case ShiftType.SAHO:
                return "bg-blue-600";
            case ShiftType.CALL:
                return "bg-green-600";
            case ShiftType.TPWE:
                return "bg-red-700";
            case ShiftType.TPHO:
                return "bg-blue-700";
            case ShiftType.FREE:
                if (free_day) return "bg-gray-200";
                return "bg-gray-100";
        }
    }

</script>

<main>
    {#if assignment.part_of_min_balance && assignment.shift_type == ShiftType.FREE}
        <div class="flex flex-none w-12 h-6 rounded-lg justify-center items-center cursor-default {free_day ? " bg-gray-200" : " bg-gray-100"}">
            <div class="text-xs text-red-500 font-bold"> ! </div>
        </div>
    {:else}
        <div class="flex flex-none w-12 h-6 {getColor(assignment.shift_type)} rounded-lg justify-center items-center cursor-default">
            <div class="text-xs font-bold text-white"> {assignment.shift_type}</div>
        </div>
    {/if}
</main>
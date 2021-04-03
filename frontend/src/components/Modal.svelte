<script context="module" lang="ts">
    const modalList: HTMLElement[] = []
</script>

<script lang="ts">
    import { booleanStore } from '../store/booleanStore'

    const store = booleanStore(false)
    const { isOpen, open, close } = store

    function keydown(e: KeyboardEvent) {
        e.stopPropagation()
            if (e.key === 'Escape') {
                close()
            }
        }

    function transitionend(e: TransitionEvent) {
        const node = e.target as HTMLElement
        node.focus()
    }

    function modalAction(node: HTMLElement) {
        const returnFn = []
        // for accessibility
        if (document.body.style.overflow !== 'hidden') {
            const original = document.body.style.overflow
            document.body.style.overflow = 'hidden'
            returnFn.push(() => {
                document.body.style.overflow = original
            })
        }
        node.addEventListener('keydown', keydown)
        node.addEventListener('transitionend', transitionend)
        node.focus()

        modalList.push(node)
        returnFn.push(() => {
            node.removeEventListener('keydown', keydown)
            node.removeEventListener('transitionend', transitionend)
            modalList.pop()
            // Optional chaining to guard against empty array.
            modalList[modalList.length - 1]?.focus()
        })
        return {
            destroy: () => returnFn.forEach((fn) => fn()),
        }
    }

</script>

<slot name="trigger" {open}>
    <!-- fallback -->
    <button on:click={open}> Open </button>
</slot>

{#if $isOpen}
<div use:modalAction tabindex="0" class="fixed top-0 left-0 w-full h-full flex justify-center items-center">
    <!-- backdrop -->
    <div on:click={close} class="fixed w-full h-full bg-white bg-opacity-60" />

    <!-- modal view -->
    <div class="z-10 px-10 py-3 border border-black rounded-md bg-white overflow-hidden">

        <div class="flex flex-row"> 
            <slot name="header" {store}>
                <!-- fallback-->
                <div>
                    Modal header goed here.
                </div>
            </slot>
        </div>

    
        <div class="max-h-3/4 max-w-7xl overflow-auto">
            <slot name="content" {store}/>
        </div>
    
        <slot name="footer" {close}>
            <!-- fallback -->
            <div>
                Modal footer goed here.
                <button on:click={close}>Close</button>
            </div>
        </slot>

    </div>
</div>
{/if}
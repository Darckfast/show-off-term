package templates

var Html = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="480" height="164"
    style="background-image: none !important;">
    <foreignObject width="480" height="164">
        <div xmlns="http://www.w3.org/1999/xhtml">
            <style>
                :before,
                :after {
                    --tw-content: '';
                }

                .main {
                    font-family: Roboto, sans-serif;
                }
                .absolute {
                    position: absolute;
                }

                .relative {
                    position: relative;
                }

                .left-0 {
                    left: 0;
                }

                .right-4 {
                    right: 1rem;
                }

                .top-0 {
                    top: 0;
                }

                .flex {
                    display: flex;
                }

                .inline-flex {
                    display: inline-flex;
                }

                .h-auto {
                    height: auto;
                }

                .w-full {
                    width: 27rem;
                }

                @keyframes pulse {
                    50% {
                        opacity: 0.5;
                    }
                }

                .animate-pulse {
                    animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
                }

                .justify-center {
                    justify-content: center;
                }

                .rounded-md {
                    border-radius: 0.375rem;
                }

                .rounded-b-none {
                    border-bottom-right-radius: 0;
                    border-bottom-left-radius: 0;
                }

                .border {
                    border-width: 1px;
                }

                .border-slate-600 {
                    --tw-border-opacity: 1;
                    border-color: rgb(71 85 105 / var(--tw-border-opacity));
                }

                .bg-slate-800 {
                    --tw-bg-opacity: 1;
                    background-color: rgb(30 41 59 / var(--tw-bg-opacity));
                }

                .bg-slate-900 {
                    --tw-bg-opacity: 1;
                    background-color: rgb(15 23 42 / var(--tw-bg-opacity));
                }

                .p-1 {
                    padding-top: 0.25rem;
                    padding-bottom: 0.25rem;
                }

                .p-4 {
                    padding: 1rem;
                }

                .pt-10 {
                    padding-top: 2.5rem;
                }

                .align-middle {
                    vertical-align: middle;
                }

                .font-semibold {
                    font-weight: 600;
                }

                .italic {
                    font-style: italic;
                }

                .text-white {
                    --tw-text-opacity: 1;
                    color: rgb(255 255 255 / var(--tw-text-opacity));
                }

                .str {
                    --tw-text-opacity: 1;
                    color: rgb(249 115 22 / var(--tw-text-opacity));
                }

                .kwd {
                    --tw-text-opacity: 1;
                    color: rgb(168 85 247 / var(--tw-text-opacity));
                }

                .com {
                    color: #32cd32;
                    font-style: italic;
                }

                .typ {
                    color: #6ecbcc;
                }

                .lit {
                    color: #d06;
                }

                .pun {
                    --tw-text-opacity: 1;
                    color: rgb(75 85 99 / var(--tw-text-opacity));
                }

                .pln {
                    --tw-text-opacity: 1;
                    color: rgb(241 245 249 / var(--tw-text-opacity));
                }

                .dec {
                    color: #3387cc;
                }

                a {
                    text-decoration: none;
                }

                .after\:h-4:after {
                    content: var(--tw-content);
                    height: 1rem;
                }

                .after\:w-2:after {
                    content: var(--tw-content);
                    width: 0.5rem;
                }

                .after\:bg-white:after {
                    content: var(--tw-content);
                    --tw-bg-opacity: 1;
                    background-color: rgb(255 255 255 / var(--tw-bg-opacity));
                }

                pre {
                    margin: 0;
                }
            </style>
            <div
                class="main border border-slate-600 h-auto w-full flex text-white bg-slate-900 relative rounded-md p-4 pt-10">
                <div style="width: 100%"
                    class="p-1 rounded-b-none flex justify-center bg-slate-800 absolute top-0 left-0 rounded-md">
                    $ {{.Lang}}
                    <div class="absolute right-4">🗕 🗖 ⨉</div>
                </div>
                <pre
                    class="w-full font-semibold">{{.Code}}<span class="inline-flex align-middle after:w-2 after:h-4 after:bg-white animate-pulse" /></pre>
            </div>
        </div>
    </foreignObject>
</svg>
`

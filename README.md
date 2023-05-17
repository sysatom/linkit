# Linkit

[![build](https://github.com/sysatom/linkit/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/sysatom/linkit/actions/workflows/build.yml)

## Getting Started

```
gh repo clone agmmnn/tauri-ui
cd tauri-ui
pnpm i
```

```
pnpm tauri dev
pnpm tauri build
```

## Update Components

Note that **shadcn/ui** [is not a library](https://ui.shadcn.com/docs#faqs), therefore you will need to update the components manually. To do so, you can [download](https://download-directory.github.io/?url=https%3A%2F%2Fgithub.com%2Fshadcn%2Fui%2Ftree%2Fmain%2Fapps%2Fwww%2Fcomponents%2Fui) the _[shadcn/ui/apps/www/components/ui](https://github.com/shadcn/ui/tree/main/apps/www/components/ui)_ directory and paste it into _[src/components/ui](/src/components/ui)_.

## Folder Structure

```
.
├── next-env.d.ts
├── next.config.js    //nextjs config file https://nextjs.org/docs/api-reference/next.config.js/introduction
├── package.json
├── postcss.config.js
├── README.md
├── public
├── src               //frontend src:
│   ├── app           //next.js appdir https://beta.nextjs.org/docs/routing/fundamentals
│   ├── assets
│   ├── components    //from shadcn/ui
│   │   └── ui
│   ├── data
│   ├── hooks
│   ├── lib
│   └── styles
├── src-tauri         //backend src:
│   ├── build.rs
│   ├── Cargo.lock
│   ├── Cargo.toml    //https://doc.rust-lang.org/cargo/reference/manifest.html
│   ├── icons         //https://tauri.app/v1/guides/features/icons/
│   ├── src           //rust codes
│   └── tauri.conf.json  //tauri config file https://next--tauri.netlify.app/next/api/config
├── prettier.config.js     //prettier config file https://prettier.io/docs/en/configuration.html
├── tailwind.config.js     //tailwind config file https://tailwindcss.com/docs/configuration
└── tsconfig.json          //typescript config file https://www.typescriptlang.org/docs/handbook/tsconfig-json.html
```

## Recommended IDE Setup

- [VS Code](https://code.visualstudio.com/) + [Tauri](https://marketplace.visualstudio.com/items?itemName=tauri-apps.tauri-vscode) + [rust-analyzer](https://marketplace.visualstudio.com/items?itemName=rust-lang.rust-analyzer)

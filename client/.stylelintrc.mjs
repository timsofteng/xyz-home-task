export default {
  extends: ["stylelint-config-standard", "stylelint-config-clean-order"],
  rules: {
    "declaration-property-value-no-unknown": true,
    "media-query-no-invalid": null,
    "selector-class-pattern": null,
    "selector-id-pattern": null,
  },
  customSyntax: "postcss-styled-syntax",
};

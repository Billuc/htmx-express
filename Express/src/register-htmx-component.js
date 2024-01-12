import express from "express";

/**
 *
 * @param {express.Request} req
 * @param {express.Response} res
 * @param {*} next
 */
export default function registerHtmxComponent(componentModule) {
  const componentName = componentModule.name;
  const componentStyle = componentModule.style;

  return function (req, res, next) {
    if (req.path == `/api/style/${componentName}.css`) {
      res.send(componentStyle);
    } else {
      next();
    }
  };
}

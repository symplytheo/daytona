/* tslint:disable */

/**
 * Daytona
 * Daytona AI platform API Docs
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@daytona.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 *
 * @export
 * @enum {string}
 */

export const SnapshotState = {
  BUILD_PENDING: 'build_pending',
  BUILDING: 'building',
  PENDING: 'pending',
  PULLING: 'pulling',
  PENDING_VALIDATION: 'pending_validation',
  VALIDATING: 'validating',
  ACTIVE: 'active',
  INACTIVE: 'inactive',
  ERROR: 'error',
  BUILD_FAILED: 'build_failed',
  REMOVING: 'removing',
} as const

export type SnapshotState = (typeof SnapshotState)[keyof typeof SnapshotState]
